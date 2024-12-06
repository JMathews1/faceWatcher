package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	// Open webcam
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println("Error opening webcam:", err)
		return
	}
	defer webcam.Close()

	// Reduce resolution for performance
	webcam.Set(gocv.VideoCaptureFrameWidth, 640)
	webcam.Set(gocv.VideoCaptureFrameHeight, 480)

	// Prepare a window to display the video feed
	window := gocv.NewWindow("Focus Tracker")
	defer window.Close()

	// Create a matrix to store frames
	frame := gocv.NewMat()
	defer frame.Close()

	// Grayscale matrix for face detection
	grayFrame := gocv.NewMat()
	defer grayFrame.Close()

	// Load a pre-trained face detection model (Haar Cascade)
	cascade := gocv.NewCascadeClassifier()
	defer cascade.Close()
	if !cascade.Load("haarcascade_frontalface_default.xml") {
		fmt.Println("Error loading cascade file")
		return
	}

	// Timer variables
	var startTime time.Time
	var totalFocusedTime time.Duration
	tracking := false

	// File for logging
	logFile, err := os.OpenFile("focus_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)

	fmt.Println("Starting Focus Tracker...")

	frameCount := 0

	for {
		// Read the current frame
		if ok := webcam.Read(&frame); !ok || frame.Empty() {
			fmt.Println("Cannot read frame")
			break
		}

		// Skip frames for reduced processing
		frameCount++
		if frameCount%5 != 0 { // Process every 5th frame
			continue
		}

		// Convert to grayscale for faster processing
		gocv.CvtColor(frame, &grayFrame, gocv.ColorBGRToGray)

		// Detect faces in the grayscale frame
		faces := cascade.DetectMultiScaleWithParams(grayFrame, 1.1, 5, 0, image.Point{X: 50, Y: 50}, image.Point{})

		if len(faces) > 0 {
			if !tracking {
				// Start tracking
				startTime = time.Now()
				tracking = true
				fmt.Println("Focus started.")
			} else {
				// Display current focus time
				currentFocus := time.Since(startTime)
				text := fmt.Sprintf("Current Focus: %s", formatDuration(currentFocus))
				gocv.PutText(&frame, text, image.Point{X: 10, Y: 30}, gocv.FontHersheySimplex, 1.0, color.RGBA{255, 255, 255, 255}, 2)
			}
		} else {
			if tracking {
				// End tracking
				focusDuration := time.Since(startTime)
				tracking = false

				// Ignore short focus durations
				if focusDuration >= 1*time.Minute {
					totalFocusedTime += focusDuration
					logger.Printf("Session Focus Time: %s\n", formatDuration(focusDuration))
					fmt.Printf("Session logged: %s\n", formatDuration(focusDuration))
				} else {
					fmt.Println("Session too short. Not logged.")
				}
			}
		}

		// Display the frame in the window
		window.IMShow(frame)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

	// Log total focus time at the end of the program
	if totalFocusedTime >= 1*time.Minute {
		logger.Printf("Total Focus Time: %s\n", formatDuration(totalFocusedTime))
		fmt.Printf("Total session logged: %s\n", formatDuration(totalFocusedTime))
	} else {
		fmt.Println("Total session too short. Not logged.")
	}
}

// Helper function to format durations into human-readable strings
func formatDuration(d time.Duration) string {
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02dh:%02dm:%02ds", h, m, s)
}