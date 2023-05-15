package com.tennis.tennisreservation.models;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.PrimaryKeyJoinColumn;

import jakarta.validation.constraints.NotEmpty;

@Entity
@PrimaryKeyJoinColumn(name = "id")
public class Member extends User {

    @Column(name="join_date")
    @NotEmpty(message = "*Please provide your join date")
    private Date joinDate;
    
    @Column(name="level")
    @NotEmpty(message = "*Please provide your level")
    private String level;

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
