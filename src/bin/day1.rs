fn part1() {
    let input = include_str!("input").split("\n");

    let mut num_prev = 0;
    let mut prev = 0;

    for depth in input {
        let mut num = 0;
        match depth.parse() {
            Ok(n) => num = n,
            Err(_) => (),
        }

        if num > prev {
            num_prev += 1
        }
        prev = num
    }

    println!("{}", num_prev)
}

fn part2() {
    let input = include_str!("input").split("\n");

    let mut num_prev = 0;
    let mut prev = 0;

    for depth in input {
        let mut num = 0;
        match depth.parse() {
            Ok(n) => num = n,
            Err(_) => (),
        }

        if num > prev {
            num_prev += 1
        }
        prev = num
    }

    println!("{}", num_prev)
}

fn main() {
    //part1();
    part2();
}
