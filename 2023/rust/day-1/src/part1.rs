use anyhow::{Error, Result};

pub fn process(input: &str) -> Result<String, Error> {
    let output = input
        .lines()
        .map(|line| {
            let mut iter = line.chars();

            let first = iter.find_map(|c| c.to_digit(10)).expect("not a number");

            let last = iter
                .rfind(|c| c.is_ascii_digit())
                .map(|c| c.to_digit(10).unwrap())
                .unwrap_or(first);
            first * 10 + last
        })
        .sum::<u32>();

    Ok(output.to_string())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_process() -> anyhow::Result<()> {
        let input = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";
        assert_eq!("142", process(input)?);
        Ok(())
    }
}
