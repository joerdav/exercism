use std::fmt::Debug;

pub struct Clock {
    minutes: i32,
}

impl Debug for Clock {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Clock")
            .field("minutes", &self.minutes)
            .finish()
    }
}

impl PartialEq for Clock {
    fn eq(&self, other: &Self) -> bool {
        self.minutes == other.minutes
    }
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let raw = minutes + hours * 60;
        let mins = raw % (24 * 60);
        Clock {
            minutes: if mins < 0 { mins + (24 * 60) } else { mins },
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(0, self.minutes + minutes)
    }

    pub fn to_string(&self) -> String {
        let h = self.minutes / 60;
        let m = self.minutes % 60;
        format!("{:02}:{:02}", h, m)
    }
}
