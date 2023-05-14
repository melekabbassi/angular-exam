package com.tennis.tennisreservation.repositories;

import org.springframework.data.jpa.repository.JpaRepository;

import com.tennis.tennisreservation.models.Admin;

public interface IAdminRepo extends JpaRepository<Admin, Integer> {
    
}
