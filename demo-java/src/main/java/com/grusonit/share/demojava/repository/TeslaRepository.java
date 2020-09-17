package com.grusonit.share.demojava.repository;

import com.grusonit.share.demojava.models.Tesla;
import org.springframework.data.jpa.repository.JpaRepository;

public interface TeslaRepository extends JpaRepository<Tesla, Long> {
}
