import org.junit.Test;

import java.io.IOException;

import static org.junit.Assert.*;

public class ConverterTest {

    Converter converter = new Converter();

    @Test
    public void convertUAHCurrencyTest() throws IOException {
        double expect = 261.07966999999996;

        double result = converter.convertCurrency(10, "UAH");

        assertEquals(expect, result, 0.0001);
    }

    @Test
    public void convertEURCurrencyTest() throws IOException {
        double expect = 8.47498;

        double result = converter.convertCurrency(10, "EUR");

        assertEquals(expect, result, 0.0001);
    }

    @Test
    public void convertGBPCurrencyTest() throws IOException {
        double expect = 7.4732;

        double result = converter.convertCurrency(10, "GBP");

        assertEquals(expect, result, 0.0001);
    }

    @Test(expected = NullPointerException.class)
    public void convertRUBCurrencyTest() throws IOException {

        double result = converter.convertCurrency(10, "RUB");

    }
}