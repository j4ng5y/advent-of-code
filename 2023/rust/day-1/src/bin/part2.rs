use day_1::part2::process;

fn main() -> Result<(), anyhow::Error> {
    let file = include_str!("../../input1.txt");
    let result = process(file)?;
    println!("Result: {}", result);
    Ok(())
}
