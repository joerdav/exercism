#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    if _first_list == _second_list {
        return Comparison::Equal;
    }
    if _first_list.len() > _second_list.len() {
        if sub(_second_list, _first_list) {
            return Comparison::Superlist;
        }
    }
    if _second_list.len() > _first_list.len() {
        if sub(_first_list, _second_list) {
            return Comparison::Sublist;
        }
    }
    Comparison::Unequal
}

fn sub<T: PartialEq>(sub: &[T], sup: &[T]) -> bool {
    sub.len() == 0 || sup.windows(sub.len()).any(|x| x == sub)
}
