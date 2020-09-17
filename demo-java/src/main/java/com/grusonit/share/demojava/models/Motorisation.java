package com.grusonit.share.demojava.models;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

@Entity
public class Motorisation {
    @Id
    @GeneratedValue
    public long id;
    public double acceleration;
    public int autonomy;
}
