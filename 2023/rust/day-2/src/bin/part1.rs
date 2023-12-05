use day_2::part1::process;

fn main() -> Result<(), anyhow::Error> {
    let file = include_str!("../../input.txt");
    let result = process(file)?;
    println!("Result: {}", result);
    Ok(())
}
