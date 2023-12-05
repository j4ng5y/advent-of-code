use day_1::part1::process;

fn main() -> Result<(), anyhow::Error> {
    let file = include_str!("../../input1.txt");
    let result = process(file).ok().unwrap();
    println!("Result: {}", result);

    Ok(())
}
