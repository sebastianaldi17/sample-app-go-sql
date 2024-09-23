package com.sebastianaldi17.sbrestapi;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.security.servlet.SecurityAutoConfiguration;

@SpringBootApplication(exclude = {SecurityAutoConfiguration.class})
public class SbrestapiApplication {

	public static void main(String[] args) {
		SpringApplication.run(SbrestapiApplication.class, args);
	}

}
