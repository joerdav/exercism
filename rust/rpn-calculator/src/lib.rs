#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut stack: Vec<i32> = Vec::new();
    for curr in inputs {
        let push = match curr {
            CalculatorInput::Value(v) => *v,
            _ => match (stack.pop(), stack.pop()) {
                (Some(r), Some(l)) => match curr {
                    CalculatorInput::Add => l + r,
                    CalculatorInput::Subtract => l - r,
                    CalculatorInput::Multiply => l * r,
                    CalculatorInput::Divide => l / r,
                    _ => panic!("not possible"),
                },
                _ => return None,
            },
        };
        stack.push(push);
    }
    if stack.len() == 1 {
        Some(stack[0])
    } else {
        None
    }
}
