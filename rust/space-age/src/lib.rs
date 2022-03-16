// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

#[derive(Debug)]
pub struct Duration {
    seconds: u64,
}

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Duration { seconds: s }
    }
}

pub trait Planet {
    fn orbital_period() -> f64;
    fn years_during(d: &Duration) -> f64 {
        return d.seconds as f64 / 31557600.0 / Self::orbital_period();
    }
}

macro_rules! years_during {
    ($name:ident, $time:expr) => {
        pub struct $name;
        impl Planet for $name {
            fn orbital_period() -> f64 {
                $time
            }
        }
    };
}

years_during!(Mercury, 0.2408467);
years_during!(Venus, 0.61519726);
years_during!(Earth, 1.0);
years_during!(Mars, 1.8808158);
years_during!(Jupiter, 11.862615);
years_during!(Saturn, 29.447498);
years_during!(Uranus, 84.016846);
years_during!(Neptune, 164.79132);
