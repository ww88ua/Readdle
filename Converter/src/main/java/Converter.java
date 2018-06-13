import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParser;

import java.io.*;
import java.net.URL;
import java.net.URLConnection;
import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;

public class Converter {

    public static void main(String[] args) throws IOException {
        Scanner scanner = new Scanner(System.in);
        System.out.println("Enter currency to convert: ");
        String toCurrency = scanner.nextLine();
        System.out.println("Enter amount to convert: ");
        int amount = scanner.nextInt();

        Converter converter = new Converter();
        double result = converter.convertCurrency(amount, toCurrency);
        System.out.println("Result: " + result + " "+toCurrency.toUpperCase());
    }

    public Double convertCurrency(double amount, String toCurrency) throws IOException {
        Set<String> allowCurrencies = new HashSet<String>();
        allowCurrencies.add("UAH");
        allowCurrencies.add("EUR");
        allowCurrencies.add("GBP");

        if (allowCurrencies.contains(toCurrency)) {

            String query = "USD_" + toCurrency.toUpperCase();
            String urlString = "https://free.currencyconverterapi.com/api/v5/convert?q=" + query + "&compact=ultra";
            URL url = new URL(urlString);
            URLConnection conn = url.openConnection();
            conn.connect();
            JsonParser jsonParser = new JsonParser();
            JsonElement root = jsonParser.parse(new InputStreamReader((InputStream) conn.getContent()));
            JsonObject jsonObject = root.getAsJsonObject();
            String value = jsonObject.get(query).getAsString();

            return amount * Double.parseDouble(value);
        }
        else {
            System.out.println("Not allow type of currency");
            return null;
        }
    }



}
