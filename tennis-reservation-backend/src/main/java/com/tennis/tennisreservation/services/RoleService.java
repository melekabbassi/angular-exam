package com.tennis.tennisreservation.services;


import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;

import com.tennis.tennisreservation.models.Role;
import com.tennis.tennisreservation.repositories.IRoleRepo;

public class RoleService {
    private final IRoleRepo roleRepo;

    @Autowired
    public RoleService(IRoleRepo roleRepo) {
        this.roleRepo = roleRepo;
    }
    
    public List<Role> findAllRoles() {
        return roleRepo.findAll();
    }

    public Role getRoleById(Long id) {
        return roleRepo.findById(id).orElseThrow(() -> new RuntimeException("Role by id " + id + " was not found"));    
    }

    public Role createRole(Role role) {
        return roleRepo.save(role);
    }

    public Role updateRole(Long id, Role updatedRole) {
        Role existingRole = getRoleById(id);
        existingRole.setRole(updatedRole.getRole());
        existingRole.setUsers(updatedRole.getUsers());
        return roleRepo.save(existingRole);        
    }

    public void deleteRole(Long id) {
        roleRepo.deleteById(id);
    }
}
