package com.example.kudaki;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.view.animation.Animation;
import android.view.animation.AnimationUtils;
import android.widget.Button;
import android.widget.TextView;

import androidx.appcompat.app.AppCompatActivity;
import androidx.viewpager.widget.ViewPager;

import com.example.kudaki.adapter.TutorialViewPagerAdapter;
import com.example.kudaki.login.LoginActivity;
import com.example.kudaki.model.TutorialItem;
import com.google.android.material.tabs.TabLayout;
import com.orhanobut.hawk.Hawk;

import java.util.ArrayList;

import butterknife.BindView;
import butterknife.ButterKnife;

public class TutorialActivity extends AppCompatActivity {
    @BindView(R.id.tutorialTab)
    TabLayout tabLayout;
    @BindView(R.id.tutorialViewPager)
    ViewPager viewPager;
    @BindView(R.id.tutorialNext)
    TextView next;
    @BindView(R.id.tutorialSkip)
    TextView skip;
    @BindView(R.id.tutorialGetstarted)
    Button getStarted;

    Animation animation;
    int position = 0;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_tutorial);
        ButterKnife.bind(this);

        animation = AnimationUtils.loadAnimation(getApplicationContext(), R.anim.get_started_animation);

        ArrayList<TutorialItem> list = new ArrayList<>();
        list.add(new TutorialItem("Temukan Acara Outdoors",
                "Cari tahu dan ikuti acara outdoors favoritmu", R.drawable.tutorial1));
        list.add(new TutorialItem("Cari Informasi Destinasimu",
                "Cari tahu informasi tentang destinasi yang akan kamu kunjungi.", R.drawable.tutorial2));
        list.add(new TutorialItem("Menyewa Peralatan",
                "Kamu dapat menyewa peralatan untuk kebutuhan mendaki disekitarmu.", R.drawable.tutorial3));
        list.add(new TutorialItem("Cari Rekomendasi Peralatan",
                "Temukan rekomendasi peralatan yang sesuai dengan tujuanmu.", R.drawable.tutorial4));

        TutorialViewPagerAdapter adapter = new TutorialViewPagerAdapter(this, list);
        viewPager.setAdapter(adapter);
        tabLayout.setupWithViewPager(viewPager);

        next.setOnClickListener(v -> {
            position = viewPager.getCurrentItem();
            if (position < list.size()) {
                position++;
                viewPager.setCurrentItem(position);
            }

            if (position == list.size()-1) {
                loadLastScreen();
            }
        });

        skip.setOnClickListener(v -> {
            position = viewPager.getCurrentItem();
            if (position < list.size()) {
                viewPager.setCurrentItem(list.size()-1);
                loadLastScreen();
            }
        });

        getStarted.setOnClickListener(v -> {
            Hawk.init(this).build();
            Hawk.put("isIntroOpened", true);

            Intent login = new Intent(v.getContext(), LoginActivity.class);
            startActivity(login);
            finish();
        });

        tabLayout.addOnTabSelectedListener(new TabLayout.BaseOnTabSelectedListener() {
            @Override
            public void onTabSelected(TabLayout.Tab tab) {
                if (tab.getPosition() == list.size()-1) {
                    loadLastScreen();
                }
            }

            @Override
            public void onTabUnselected(TabLayout.Tab tab) {

            }

            @Override
            public void onTabReselected(TabLayout.Tab tab) {

            }
        });
    }

    void loadLastScreen(){
        getStarted.setVisibility(View.VISIBLE);
        getStarted.setAnimation(animation);

        tabLayout.setVisibility(View.INVISIBLE);
        next.setVisibility(View.INVISIBLE);
        skip.setVisibility(View.INVISIBLE);
    }
}