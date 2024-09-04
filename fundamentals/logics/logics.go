package main

func shopping(firstJob, secondJob bool) (bool, bool, bool) {
	buyTV50 := firstJob && secondJob;
	buyTV32 := firstJob != secondJob;
	buyIceCream := firstJob || secondJob;

	return buyTV50, buyTV32, buyIceCream;
}

func main() {
	tv50, tv32, iceCream := shopping(true, true);
	println(tv50, tv32, iceCream);

	tv50, tv32, iceCream = shopping(true, false);
	println(tv50, tv32, iceCream);

	tv50, tv32, iceCream = shopping(false, true);
	println(tv50, tv32, iceCream);

	tv50, tv32, iceCream = shopping(false, false);
	println(tv50, tv32, iceCream);
}
