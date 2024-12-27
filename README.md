# Hydration Reminder

Hydration Reminder is a simple terminal-based application that helps you stay hydrated throughout the day. By setting up work hours, lunch breaks, and your weight, this app will send periodic reminders to drink water during your active hours.

## Features

Calculates daily water requirement based on weight.
Allows configuration of work hours and lunch breaks.
Sends reminders at regular intervals based on user-defined settings.

## Installation

Clone the repository:
~~~bash
git clone https://github.com/yourusername/hydration-reminder.git
cd hydration-reminder
~~~

Build the application:
~~~bash
make build
~~~

## Configuration

The application requires a config.yaml file in the same directory as the binary. This file contains the necessary settings for the application. Below is an example configuration:

~~~yaml
weight: 75                    # Your weight in kilograms
office_hours: 8               # Total work hours in a day
lunch_interval_minutes: 60    # Duration of the lunch break in minutes
lunch_interval_start: 12      # Hour when lunch break starts (24-hour format)
~~~

## Usage

Run the application:
~~~bash
./hydration-reminder
~~~

## Notes
The application ensures that reminders are only sent during work hours and never beyond 18:00.
If the config.yaml file is missing or incorrectly configured, the application will not start.
