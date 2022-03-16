use std::{collections::HashMap, thread};

pub fn frequency(input: &[&str], worker_count: usize) -> HashMap<char, usize> {
    if input.is_empty() {
        return HashMap::new();
    }
    let num_batches = ((input.len() as f64) / (worker_count as f64)).ceil() as usize;
    input
        .chunks(num_batches)
        .map(|c| {
            let ch: Vec<String> = c.iter().map(|s| s.to_string()).collect();
            thread::spawn(move || {
                ch.iter().map(|s| s.to_lowercase()).fold(
                    HashMap::<char, usize>::new(),
                    |mut acc, s| {
                        let vals = s.chars().fold(HashMap::<char, usize>::new(), |mut acc, w| {
                            if !w.is_alphabetic() {
                                return acc;
                            }
                            *acc.entry(w).or_insert(0) += 1;
                            acc
                        });

                        for (key, val) in vals {
                            *acc.entry(key).or_insert(0) += val;
                        }
                        acc
                    },
                )
            })
        })
        .map(|h| h.join().unwrap())
        .fold(HashMap::<char, usize>::new(), |mut acc, kv| {
            for (key, val) in kv {
                *acc.entry(key).or_insert(0) += val;
            }
            acc
        })
}
