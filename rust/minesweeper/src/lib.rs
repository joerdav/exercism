fn tile_to_num(board: &[String], x: i8, y: i8) -> u32 {
    if x < 0 || y < 0 {
        return 0;
    }
    board
        .iter()
        .nth(y as usize)
        .and_then(|r| r.chars().nth(x as usize))
        .and_then(|l| match l {
            '*' => Some(1),
            _ => Some(0),
        })
        .unwrap_or(0)
}
pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let board: Vec<String> = minefield.iter().map(|&x| x.to_string().clone()).collect();
    minefield
        .iter()
        .enumerate()
        .map(|(yc, &s)| {
            s.chars()
                .enumerate()
                .map(|(xc, c)| match c {
                    '*' => '*',
                    _ => {
                        let x = xc as i8;
                        let y = yc as i8;
                        let n = tile_to_num(&board, x - 1, y - 1)
                            + tile_to_num(&board, x, y - 1)
                            + tile_to_num(&board, x + 1, y - 1)
                            + tile_to_num(&board, x - 1, y)
                            + tile_to_num(&board, x + 1, y)
                            + tile_to_num(&board, x - 1, y + 1)
                            + tile_to_num(&board, x, y + 1)
                            + tile_to_num(&board, x + 1, y + 1);
                        match n.to_string().chars().nth(0).unwrap() {
                            '0' => ' ',
                            a => a,
                        }
                    }
                })
                .collect::<String>()
        })
        .collect()
}
