use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    input
        .graphemes(true)
        .fold("".to_string(), |acc, c| c.to_string() + &acc)
}
