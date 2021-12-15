// This stub file contains items which aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut map = magazine
        .iter()
        .fold(HashMap::<String, u16>::new(), |mut acc, n| {
            *acc.entry(n.to_string()).or_insert(0) += 1;
            acc
        });
    for word in note {
        let count = map.entry(word.to_string()).or_default();
        if *count == 0 {
            return false;
        }
        *map.entry(word.to_string()).or_default() -= 1;
    }

    true
}
