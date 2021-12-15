#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut s = vec![];
    for inp in inputs {
        match inp {
            &CalculatorInput::Add => {
                let b = s.pop();
                let a = s.pop();
                if a.is_none() || b.is_none() {
                    return None;
                }
                s.push(a.unwrap() + b.unwrap())
            }
            &CalculatorInput::Subtract => {
                let b = s.pop();
                let a = s.pop();
                if a.is_none() || b.is_none() {
                    return None;
                }
                s.push(a.unwrap() - b.unwrap())
            }
            &CalculatorInput::Multiply => {
                let b = s.pop();
                let a = s.pop();
                if a.is_none() || b.is_none() {
                    return None;
                }
                s.push(a.unwrap() * b.unwrap())
            }
            &CalculatorInput::Divide => {
                let b = s.pop();
                let a = s.pop();
                if a.is_none() || b.is_none() {
                    return None;
                }
                s.push(a.unwrap() / b.unwrap())
            }
            &CalculatorInput::Value(v) => s.push(v),
        }
    }
    if s.len() != 1 {
        return None;
    }
    Some(s[0])
}
