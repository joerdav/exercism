use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock(i32);

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mins = minutes + hours * 60;
        Clock(mins.rem_euclid(24 * 60))
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(0, self.0 + minutes)
    }

    fn hours(&self) -> i32 {
        self.0 / 60
    }
    fn minutes(&self) -> i32 {
        self.0 % 60
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours(), self.minutes())
    }
}
