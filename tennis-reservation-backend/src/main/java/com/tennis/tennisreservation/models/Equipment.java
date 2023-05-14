package com.tennis.tennisreservation.models;

public class Equipment {
    private int id;
    private String type;
    private Boolean isAvailable;
    private double price;

    public Equipment(int id, String type, Boolean isAvailable, double price) {
        this.id = id;
        this.type = type;
        this.isAvailable = isAvailable;
        this.price = price;
    }

    public Equipment(String type, Boolean isAvailable, double price) {
        this.type = type;
        this.isAvailable = isAvailable;
        this.price = price;
    }

    public Equipment() {}

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Boolean getIsAvailable() {
        return isAvailable;
    }

    public void setIsAvailable(Boolean isAvailable) {
        this.isAvailable = isAvailable;
    }

    public double getPrice() {
        return price;
    }

    public void setPrice(double price) {
        this.price = price;
    }
    
}
