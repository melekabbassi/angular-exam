package com.tennis.tennisreservation.models;

import java.util.Date;
import java.util.List;

public class Reservation {

    private int id;
    private int userId;
    private int courtId;
    private String type;
    private Date date;
    private int hour;
    private int duration;
    private List<String> equipment;
    private List<String> services;

    public Reservation(int id, int userId, int courtId, String type, Date date, int hour, int duration, List<String> equipment, List<String> services) {
        this.id = id;
        this.userId = userId;
        this.courtId = courtId;
        this.date = date;
        this.hour = hour;
        this.duration = duration;
        this.equipment = equipment;
        this.services = services;
    }

    public Reservation(int userId, int courtId, String type, Date date, int hour, int duration, List<String> equipment, List<String> services) {
        this.userId = userId;
        this.courtId = courtId;
        this.date = date;
        this.hour = hour;
        this.duration = duration;
        this.equipment = equipment;
        this.services = services;
    }

    public Reservation() {}

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getUserId() {
        return userId;
    }

    public void setUserId(int userId) {
        this.userId = userId;
    }

    public int getCourtId() {
        return courtId;
    }

    public void setCourtId(int courtId) {
        this.courtId = courtId;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Date getDate() {
        return date;
    }

    public void setDate(Date date) {
        this.date = date;
    }

    public int getHour() {
        return hour;
    }

    public void setHour(int hour) {
        this.hour = hour;
    }

    public int getDuration() {
        return duration;
    }

    public void setDuration(int duration) {
        this.duration = duration;
    }

    public List<String> getEquipment() {
        return equipment;
    }

    public void setEquipment(List<String> equipment) {
        this.equipment = equipment;
    }

    public List<String> getServices() {
        return services;
    }

    public void setServices(List<String> services) {
        this.services = services;
    }
    
}
