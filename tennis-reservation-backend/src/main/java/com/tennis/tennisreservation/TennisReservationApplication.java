package com.tennis.tennisreservation;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@EntityScan("com.tennis.tennisreservation.models")
@ComponentScan(basePackages = {"com.tennis.tennisreservation.controllers", "com.tennis.tennisreservation.services", "com.tennis.tennisreservation.config"})
public class TennisReservationApplication {

	public static void main(String[] args) {
		SpringApplication.run(TennisReservationApplication.class, args);
	}

}
