package com.tennis.tennisreservation.services;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import com.tennis.tennisreservation.models.Admin;
import com.tennis.tennisreservation.repositories.IAdminRepo;

@Service
public class AdminService {
    private final IAdminRepo adminRepo;
    private final BCryptPasswordEncoder passwordEncoder;

    @Autowired
    public AdminService(IAdminRepo adminRepo, BCryptPasswordEncoder passwordEncoder) {
        this.adminRepo = adminRepo;
        this.passwordEncoder = passwordEncoder;
    }

    public List<Admin> findAllAdmins() {
        return adminRepo.findAll();
    }

    public Admin getAdminById(Long id) {
        return adminRepo.findById(id).orElseThrow(() -> new RuntimeException("Admin by id " + id + " was not found"));    
    }

    public Admin createAdmin(Admin admin) {
        admin.setPassword(passwordEncoder.encode(admin.getPassword()));
        return adminRepo.save(admin);
    }

    public Admin updateAdmin(Long id, Admin updatedAdmin) {
        Admin existingAdmin = getAdminById(id);
        existingAdmin.setFirstName(updatedAdmin.getFirstName());
        existingAdmin.setLastName(updatedAdmin.getLastName());
        existingAdmin.setEmail(updatedAdmin.getEmail());
        existingAdmin.setStatus(updatedAdmin.getStatus());
        existingAdmin.setIsActive(updatedAdmin.getIsActive());
        existingAdmin.setRoles(updatedAdmin.getRoles());
        return adminRepo.save(existingAdmin);        
    }

    public void deleteAdmin(Long id) {
        adminRepo.deleteById(id);
    }

}
