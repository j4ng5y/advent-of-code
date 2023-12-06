use day_6::part1::process;

fn main() -> anyhow::Result<()> {
    let file = include_str!("../../input.txt");
    let result = process(file)?;
    println!("{}", result);
    Ok(())
}
