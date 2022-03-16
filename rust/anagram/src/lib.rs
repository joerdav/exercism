use std::collections::HashSet;

fn sort(s: &str) -> String {
    let mut l = s.to_lowercase().chars().collect::<Vec<char>>();
    l.sort_unstable_by(|a, b| a.cmp(b));
    l.into_iter().collect()
}

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let wo = word.to_string().to_lowercase();
    let ws: String = sort(word);
    possible_anagrams
        .into_iter()
        .filter(|w| w.to_string().to_lowercase() != wo && ws == sort(w))
        .map(|&w| w)
        .collect()
}
