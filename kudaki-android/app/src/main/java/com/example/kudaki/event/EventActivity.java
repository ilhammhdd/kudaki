package com.example.kudaki.event;

import android.app.SearchManager;
import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.view.Menu;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.SearchView;
import androidx.appcompat.widget.Toolbar;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.kudaki.MainActivity;
import com.example.kudaki.R;
import com.example.kudaki.explore.ExploreActivity;
import com.example.kudaki.model.event.Event;
import com.example.kudaki.profile.ProfileActivity;
import com.example.kudaki.renting.RentalActivity;
import com.google.android.material.bottomnavigation.BottomNavigationView;

import java.util.ArrayList;
import java.util.List;

import butterknife.BindView;
import butterknife.ButterKnife;

public class EventActivity extends AppCompatActivity {
    @BindView(R.id.eventNav)
    BottomNavigationView bottomNav;
    @BindView(R.id.rvEvents)
    RecyclerView rvEvents;
    @BindView(R.id.eventToolbar)
    Toolbar toolbar;

    private List<Event> eventList;
    private EventAdapter eventAdapter;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_event);
        overridePendingTransition(android.R.anim.fade_in, android.R.anim.fade_out);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);

        eventList = new ArrayList<>();
        eventAdapter = new EventAdapter(this, eventList);
        rvEvents.setLayoutManager(new LinearLayoutManager(this));
        rvEvents.setAdapter(eventAdapter);
        loadEvents();
    }

    @Override
    protected void onResume() {
        super.onResume();

        bottomNav.getMenu().getItem(1).setChecked(true);
        bottomNav.setOnNavigationItemSelectedListener(menuItem -> {
            switch (menuItem.getItemId()) {
                case R.id.navHome:
                    startActivity(new Intent(this, MainActivity.class));
                    finish();
                    return true;
                case R.id.navEvent:
                    return true;
                case R.id.navExplore:
                    startActivity(new Intent(this, ExploreActivity.class));
                    finish();
                    return true;
                case R.id.navRental:
                    startActivity(new Intent(this, RentalActivity.class));
                    finish();
                    return true;
                case R.id.navProfile:
                    startActivity(new Intent(this, ProfileActivity.class));
                    finish();
                    return true;
            }
            return false;
        });
    }

    @Override
    public boolean onCreateOptionsMenu(@NonNull Menu menu) {
        getMenuInflater().inflate(R.menu.event_menu, menu);
        // Get the SearchView and set the searchable configuration
        SearchManager searchManager = (SearchManager) this.getSystemService(Context.SEARCH_SERVICE);
        SearchView searchView = (SearchView) menu.findItem(R.id.optSearchEvent).getActionView();
        // Assumes current activity is the searchable activity
        searchView.setSearchableInfo(searchManager.getSearchableInfo(this.getComponentName()));
        searchView.setIconifiedByDefault(false); // Do not iconify the widget; expand it by default
        return super.onCreateOptionsMenu(menu);
    }

    // dummy event loader
    private void loadEvents() {
        eventList.clear();
        eventList.add(new Event("Jambore 2011",
                "https://images.unsplash.com/photo-1553362200-d2f027c173e0?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"));
        eventList.add(new Event("Jambore 2012",
                "https://images.unsplash.com/photo-1553362200-d2f027c173e0?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"));
        eventList.add(new Event("Jambore 2013",
                "https://images.unsplash.com/photo-1553362200-d2f027c173e0?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"));
        eventList.add(new Event("Jambore 2014",
                "https://images.unsplash.com/photo-1553362200-d2f027c173e0?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"));
        eventList.add(new Event("Jambore 2015",
                "https://images.unsplash.com/photo-1553362200-d2f027c173e0?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"));
        eventList.add(new Event("Jambore 2016",
                "https://images.unsplash.com/photo-1553362200-d2f027c173e0?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"));
        eventList.add(new Event("Jambore 2017",
                "https://images.unsplash.com/photo-1553362200-d2f027c173e0?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"));
        eventAdapter.notifyDataSetChanged();
    }
}
