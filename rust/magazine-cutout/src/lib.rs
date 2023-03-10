// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

fn count_words(words: Vec<&str>) -> HashMap<&str, usize> {
    words.iter().fold(HashMap::new(), |mut acc, curr| {
        *acc.entry(*curr).or_insert(0) += 1;
        acc
    })
}
pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mag_count = count_words(magazine.to_vec());
    let not_count = count_words(note.to_vec());
    not_count
        .iter()
        .all(|(w, c)| mag_count.get(*w).unwrap_or(&0) >= c)
}
