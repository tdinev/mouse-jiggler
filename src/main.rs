use std::{thread, time::Duration};
use windows::Win32::Foundation::POINT;
use windows::Win32::UI::WindowsAndMessaging::{GetCursorPos, SetCursorPos};

fn main() {
    println!("Mouse jiggler started. Moving cursor back and forth by 1 pixel every minute...");

    let mut forward = true;
    let duration = Duration::from_secs(60);

    loop {
        let mut point = POINT { x: 0, y: 0 };

        unsafe {
            if GetCursorPos(&mut point).as_bool() {
                // Determine delta based on the current direction
                let delta = if forward { 1 } else { -1 };

                // Apply movement
                SetCursorPos(point.x + delta, point.y + delta);

                // Flip direction for next time
                forward = !forward;
            }
        }

        // Wait for 60 seconds before the next jiggle
        thread::sleep(duration);
    }
}
