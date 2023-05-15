package com.tennis.tennisreservation;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan({"com.tennis.tennisreservation.models", "com.tennis.tennisreservation.repositories"})
public class TennisReservationApplication {

	public static void main(String[] args) {
		SpringApplication.run(TennisReservationApplication.class, args);
	}

}
