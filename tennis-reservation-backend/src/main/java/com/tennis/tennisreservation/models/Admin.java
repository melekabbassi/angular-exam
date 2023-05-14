package com.tennis.tennisreservation.models;

import javax.persistence.Entity;
import javax.persistence.PrimaryKeyJoinColumn;

@Entity
@PrimaryKeyJoinColumn(name = "id")
public class Admin extends User {

    public Admin(int id, String lastName, String firstName, String email, String password, String status) {
        super(id, lastName, firstName, email, password, status);
    }

    public Admin(String lastName, String firstName, String email, String password, String status) {
        super(lastName, firstName, email, password, status);
    }

    public Admin(String email, String password) {
        super(email, password);
    }

    public Admin() {
        super();
    }
    
}
