package com.tennis.tennisreservation.models;

public class Court {
    
    private int id;
    private Boolean isAvailable;

    public Court(int id, Boolean isAvailable) {
        this.id = id;
        this.isAvailable = isAvailable;
    }

    public Court(Boolean isAvailable) {
        this.isAvailable = isAvailable;
    }

    public Court() {}

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public Boolean getIsAvailable() {
        return isAvailable;
    }

    public void setIsAvailable(Boolean isAvailable) {
        this.isAvailable = isAvailable;
    }

}