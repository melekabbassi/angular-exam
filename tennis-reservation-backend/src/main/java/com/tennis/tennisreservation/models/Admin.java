package com.tennis.tennisreservation.models;

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
