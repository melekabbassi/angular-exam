package com.tennis.tennisreservation.models;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.PrimaryKeyJoinColumn;

@Entity
@PrimaryKeyJoinColumn(name = "id")
public class Member extends User {

    @Column(name="join_date")
    private Date joinDate;
    
    @Column(name="level")
    private String level;

    public Member(int id, String lastName, String firstName, String email, String password, String status, Date joinDate, String level) {
        super(id, lastName, firstName, email, password, status);
        this.joinDate = joinDate;
        this.level = level;
    }

    public Member(String lastName, String firstName, String email, String password, String status, Date joinDate, String level) {
        super(lastName, firstName, email, password, status);
        this.joinDate = joinDate;
        this.level = level;
    }

    public Member(String email, String password) {
        super(email, password);
    }

    public Member() {
        super();
    }

    public Date getJoinDate() {
        return joinDate;
    }

    public void setJoinDate(Date joinDate) {
        this.joinDate = joinDate;
    }

    public String getLevel() {
        return level;
    }

    public void setLevel(String level) {
        this.level = level;
    }
        
}
