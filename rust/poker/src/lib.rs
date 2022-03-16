use std::{char, collections::HashMap};

struct Card {
    suit: char,
    rank: usize,
}
struct Hand {
    cards: Vec<Card>,
}
impl Hand {
    fn num(&self) -> usize {
        let ranks = self
            .cards
            .iter()
            .fold(HashMap::<usize, usize>::new(), |mut acc, c| {
                *acc.entry(c.rank).or_insert(0) += 1;
                acc
            });
        let foac = ranks.iter().any(|(_, &v)| v == 4);
        let toac = ranks.iter().any(|(_, &v)| v == 3);

        if foac {
            return 3;
        }
        if toac {
            return 7;
        }
        let pairs = ranks
            .iter()
            .fold(0, |acc, r| if *r.1 == 2 { acc + 1 } else { acc });
        if pairs == 2 {
            return 8;
        }
        if pairs == 1 {
            return 9;
        }
        10
    }
}

pub fn winning_hands<'a>(hands: &[&'a str]) -> Vec<&'a str> {
    unimplemented!()
}
