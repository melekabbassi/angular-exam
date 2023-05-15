package com.tennis.tennisreservation.models;

import javax.persistence.Entity;
import javax.persistence.PrimaryKeyJoinColumn;

@Entity
@PrimaryKeyJoinColumn(name = "id")
public class Admin extends User {

    public Admin(String email, String password) {
        super(email, password);
    }
    
}
