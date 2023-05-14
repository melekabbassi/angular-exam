package com.tennis.tennisreservation.models;

import java.util.Date;
import java.util.List;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToMany;
import javax.persistence.ManyToOne;

@Entity
public class Reservation {

    @Id
    @GeneratedValue(strategy= GenerationType.IDENTITY)
    @Column(name = "id", updatable = false, nullable = false)
    private int id;

    @ManyToOne
    @JoinColumn(name = "user_id")
    private User user;
    
    @ManyToOne
    @JoinColumn(name = "court_id")
    private Court court;

    @Column(name="type")
    private String type;

    @Column(name="date")
    private Date date;
    
    @Column(name="hour")
    private int hour;
    
    @Column(name="duration")
    private int duration;
    
    @ManyToMany
    private List<Equipment> equipment;
    
    @ManyToMany
    private List<Service> services;

    public Reservation(int id, User user, Court court, String type, Date date, int hour, int duration,
            List<Equipment> equipment, List<Service> services) {
        this.id = id;
        this.user = user;
        this.court = court;
        this.type = type;
        this.date = date;
        this.hour = hour;
        this.duration = duration;
        this.equipment = equipment;
        this.services = services;
    }
    
    public Reservation(User user, Court court, String type, Date date, int hour, int duration,
            List<Equipment> equipment, List<Service> services) {
        this.user = user;
        this.court = court;
        this.type = type;
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

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    public Court getCourt() {
        return court;
    }

    public void setCourt(Court court) {
        this.court = court;
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

    public List<Equipment> getEquipment() {
        return equipment;
    }

    public void setEquipment(List<Equipment> equipment) {
        this.equipment = equipment;
    }

    public List<Service> getServices() {
        return services;
    }

    public void setServices(List<Service> services) {
        this.services = services;
    }
    
}
