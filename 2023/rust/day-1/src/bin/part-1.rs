use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    if let Ok(lines) = read_lines("input_1.txt") {
        for line in lines {
            if let Ok(value) = line {
                println!("{}", value)
            }
        }
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn calculate_calibration_value(s: String) -> u16 {
    let mut calibration_value: u16 = 0;
    for c in s.chars() {
        let value: u16 = c.to_digit(10).unwrap() as u16;
    }
    calibration_value
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_main() {
        main();
    }
}
