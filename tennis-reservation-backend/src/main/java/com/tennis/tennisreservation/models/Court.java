package com.tennis.tennisreservation.models;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

@Entity
public class Court {
    @Id
    @GeneratedValue(strategy= GenerationType.IDENTITY)
    @Column(name = "id", updatable = false, nullable = false)
    private int id;

    @Column(name="is_available")
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