package main.structures;

/**
 * Created by doelia on 02/11/2015.
 */
public class Creneau {
    public String UE;
    public String location;
    public String description;
    
    @Override
    public String toString() {
    	return UE + ", location: " + location;
    }
}