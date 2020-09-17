package com.grusonit.share.demojava.controllers;

import com.grusonit.share.demojava.models.Tesla;
import com.grusonit.share.demojava.repository.TeslaRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("tesla")
public class TeslaController {
    @Autowired
    private TeslaRepository teslaRepository;

    @GetMapping("/all")
    public List<Tesla> getAll() {
        return teslaRepository.findAll();
    }
}
