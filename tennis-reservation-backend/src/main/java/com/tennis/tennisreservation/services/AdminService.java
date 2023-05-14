package com.tennis.tennisreservation.services;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.tennis.tennisreservation.models.Admin;
import com.tennis.tennisreservation.repositories.IAdminRepo;

@Service
public class AdminService {
    private final IAdminRepo adminRepo;

    @Autowired
    public AdminService(IAdminRepo adminRepo) {
        this.adminRepo = adminRepo;
    }

    public List<Admin> findAllAdmins() {
        return adminRepo.findAll();
    }

    public Admin findAdminById(int id) {
        return adminRepo.findById(id).orElseThrow(() -> new RuntimeException("Admin by id " + id + " was not found"));
    }

    public Admin createAdmin(Admin admin) {
        return adminRepo.save(admin);
    }

    public Admin updateAdmin(Admin admin) {
        return adminRepo.save(admin);
    }

    public void deleteAdmin(int id) {
        adminRepo.deleteById(id);
    }
}
