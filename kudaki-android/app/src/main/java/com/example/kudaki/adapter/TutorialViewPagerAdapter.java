package com.example.kudaki.adapter;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.viewpager.widget.PagerAdapter;

import com.example.kudaki.R;
import com.example.kudaki.model.TutorialItem;

import java.util.ArrayList;

public class TutorialViewPagerAdapter extends PagerAdapter {
    Context context;
    ArrayList<TutorialItem> items;

    public TutorialViewPagerAdapter(Context context, ArrayList<TutorialItem> items) {
        this.context = context;
        this.items = items;
    }

    @NonNull
    @Override
    public Object instantiateItem(@NonNull ViewGroup container, int position) {
        View view = LayoutInflater.from(context).inflate(R.layout.tutorial_item, null);

        ImageView image = view.findViewById(R.id.tutorialImage);
        TextView title = view.findViewById(R.id.tutorialTitle);
        TextView desc = view.findViewById(R.id.tutorialDesc);

        title.setText(items.get(position).getTitle());
        desc.setText(items.get(position).getDesc());
        image.setImageResource(items.get(position).getImage());

        container.addView(view);

        return view;
    }

    @Override
    public int getCount() {
        return items.size();
    }

    @Override
    public boolean isViewFromObject(@NonNull View view, @NonNull Object object) {
        return view == object;
    }

    @Override
    public void destroyItem(@NonNull ViewGroup container, int position, @NonNull Object object) {
        container.removeView((View) object);
    }
}
