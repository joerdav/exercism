/// Create an empty vector
pub fn create_empty() -> Vec<u8> {
    vec![]
}

/// Create a buffer of `count` zeroes.
///
/// Applications often use buffers when serializing data to send over the network.
pub fn create_buffer(count: usize) -> Vec<u8> {
    vec![0; count]
}

/// Create a vector containing the first five elements of the Fibonacci sequence.
///
/// Fibonacci's sequence is the list of numbers where the next number is a sum of the previous two.
/// Its first five elements are `1, 1, 2, 3, 5`.
pub fn fibonacci() -> Vec<u8> {
    create_buffer(5).iter().fold(vec![], |mut acc, _| {
        let a = *acc.last().unwrap_or(&1);
        let idx = acc.len().checked_sub(2);
        match idx {
            Some(oidx) => {
                acc.push(a + *acc.get(oidx).unwrap_or(&0));
            }
            None => {
                acc.push(a);
            }
        }
        acc
    })
}
