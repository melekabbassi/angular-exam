package com.tennis.tennisreservation.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.tennis.tennisreservation.models.Admin;

@Repository
public interface IAdminRepo extends JpaRepository<Admin, Long> { }
