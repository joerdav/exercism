/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let ds = code.chars().filter(|x| !x.is_whitespace()).rev();
    if ds.clone().any(|x| !x.is_numeric()) {
        return false;
    }
    if ds.clone().collect::<String>().len() < 2 {
        return false;
    }
    ds.map(|x| x.to_digit(10).unwrap())
        .enumerate()
        .map(|(i, x)| if i % 2 != 0 { x * 2 } else { x })
        .map(|x| if x > 9 { x - 9 } else { x })
        .sum::<u32>()
        % 10
        == 0
}
