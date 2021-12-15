// This stub file contains items which aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

const SPEED_PER_HOUR: f64 = 221.0;

fn get_success_rate(speed: u8) -> f64 {
    match speed {
        5..=8 => return 0.9,
        9..=10 => return 0.77,
        _ => return 1.0,
    }
}

pub fn production_rate_per_hour(speed: u8) -> f64 {
    return get_success_rate(speed) * speed as f64 * SPEED_PER_HOUR;
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    return (production_rate_per_hour(speed) / 60.0).floor() as u32;
}
