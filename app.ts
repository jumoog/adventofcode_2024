async function day1() {
    let total = 0;
    const input = Bun.file("day1.txt");
    const [left, right] = (await input.text())
        .trim()
        .split('\n')
        .reduce((acc, line) => {
            const [l, r] = line.split(/\s+/);
            acc[0].push(Number(l));
            acc[1].push(Number(r));
            return acc;
        }, [[], []] as number[][])
        .map(arr => arr.sort((a, b) => a - b));

    for (let index = 0; index < left.length; index++) {
        const valueL = left[index];
        const valueR = right[index];
        total += Math.abs(valueR - valueL);
    }
    console.log(total)
}

day1();