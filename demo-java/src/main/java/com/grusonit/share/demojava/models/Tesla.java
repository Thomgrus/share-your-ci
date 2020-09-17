package com.grusonit.share.demojava.models;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.OneToOne;

@Entity
public class Tesla {
    @Id
    @GeneratedValue
    public long id;
    @OneToOne
    public ModelTesla modelTesla;
    @OneToOne
    public Motorisation motorisation;
    public int price;
}
