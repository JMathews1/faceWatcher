# Focus Tracker

**Focus Tracker** is a Go-based application that uses your webcam to monitor your focus and activity. 

It detects your face and logs focus sessions, ignoring short distractions or movements away from your screen. The app is lightweight and optimized for low memory and CPU usage, making it ideal for locally-run, free personal productivity tracking.

---

## Features

- **Real-Time Focus Tracking**:
  - Detects when you are present and stops tracking when you move away.
  - Displays current focus duration and total focus time on the video feed.
  
- **Focus Session Logging**:
  - Logs focus sessions longer than 5 minutes to a file named `focus_log.txt`. This file logging system will be replaced with db integration & data analysis in a later release.
  
- **Optimized for Performance**:
  - Uses grayscale frames, reduced resolution, and frame skipping to lower memory and CPU usage.

---

## Prerequisites

### Software Requirements
1. **Go**:
   - [Install Go](https://go.dev/doc/install) if it’s not already on your system.
   - Verify the installation:
     ```bash
     go version
     ```

2. **OpenCV**:
   - Install OpenCV via Homebrew:
     ```bash
     brew install opencv
     ```

3. **GoCV**:
   - Install the Go bindings for OpenCV:
     ```bash
     go get -u -d gocv.io/x/gocv
     ```

4. **Haar Cascade File**:
   - Download the `haarcascade_frontalface_default.xml` file from the [OpenCV GitHub repository](https://github.com/opencv/opencv/tree/master/data/haarcascades).
   - Place it in the same directory as the `main.go` file.

---

///// will add the rest soon 
