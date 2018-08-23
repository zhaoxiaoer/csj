package c7;

import java.math.BigInteger;

public class Primes {
	private static final BigInteger ZERO = BigInteger.ZERO;
	private static final BigInteger ONE = BigInteger.ONE;
	private static final BigInteger TWO = BigInteger.TWO;
	
	private static final int ERR_VAL = 100;
	
	public static BigInteger nextPrime(BigInteger start) {
		if (isEven(start))
			start = start.add(ONE);
		else
			start = start.add(TWO);
		if (start.isProbablePrime(ERR_VAL))
			return start;
		else
			return nextPrime(start);
	}
	
	private static boolean isEven(BigInteger n) {
		return n.mod(TWO).equals(ZERO);
	}
	
	private static StringBuffer[] digits =
		{
				new StringBuffer("0"), new StringBuffer("1"),
				new StringBuffer("2"), new StringBuffer("3"),
				new StringBuffer("4"), new StringBuffer("5"),
				new StringBuffer("6"), new StringBuffer("7"),
				new StringBuffer("8"), new StringBuffer("9")
		};
	
	private static StringBuffer randomDigit(boolean isZeroOK) {
		int index;
		if (isZeroOK) {
			index = (int) Math.floor(Math.random() * 10);
		} else {
			index = 1 + (int)Math.floor(Math.random() * 9);
		}
		return digits[index];
	}
	
	public static BigInteger random(int numDigits) {
		StringBuffer s = new StringBuffer("");
		for (int i = 0; i < numDigits; i++) {
			if (i == 0) {
				s.append(randomDigit(false));
			} else {
				s.append(randomDigit(true));
			}
		}
		return new BigInteger(s.toString());
	}
}
