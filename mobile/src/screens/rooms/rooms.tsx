import React, { useState } from "react";
import {
  View,
  Text,
  StyleSheet,
  Dimensions,
  TouchableOpacity,
  FlatList,
} from "react-native";
import { Header } from "@/components/header";
import CalendarStrip from "react-native-calendar-strip";
import moment, { Moment } from "moment";
import Ionicons from "react-native-vector-icons/Ionicons";
import Octicons from "react-native-vector-icons/Octicons";

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const Rooms = () => {
  const [selectedDate, setSelectedDate] = useState<Moment>(moment());

  const handleDateSelected = (date: Moment) => {
    setSelectedDate(date);
  };

  // const events = [
  //   {
  //     id: 1,
  //     name: 'Nguyen Van A',
  //     mssv: 'MSSV: 2012345',
  //     status: "checkin",
  //     time: "06:30 am",
  //   },
  //   {
  //     id: 2,
  //     name: 'Nguyen Van B',
  //     mssv: 'MSSV: 2012989',
  //     status: "checkin",
  //     time: "06:30 pm",
  //   },
  // ];

  return (
    <View style={styles.container}>
      <View style={styles.header_container}>
        <Header title="Rooms" />
      </View>
      <View style={styles.content_container}>
        <View style={styles.calendarContainer}>
          <CalendarStrip
            style={styles.calendarStripStyle}
            showMonth={false}
            scrollable={true}
            scrollerPaging={true}
            // selectedDate={selectedDate}
            // onDateSelected={handleDateSelected}
            // calendarAnimation={{ type: "sequence", duration: 30 }}
            // daySelectionAnimation={{
            //   type: "border",
            //   borderWidth: 1,
            //   borderHighlightColor: "black",
            //   duration: 30,
            // }}
          />
        </View>
        <View style={styles.statusContainer}>
          <View style={styles.statusStyle}>
            <View style={styles.checkIn}>
              <View style={styles.checkNum}>
                <Octicons
                  name="people"
                  size={30}
                  color="rgba(52, 168, 83, 0.93)"
                />
                <Text style={styles.regular25In}>25</Text>
              </View>
              <Text style={styles.regular14}>check in</Text>
            </View>
            <View style={styles.verticalLine}></View>
            <View style={styles.checkOut}>
              <View style={styles.checkNum}>
                <Octicons
                  name="people"
                  size={30}
                  color="rgba(237, 74, 74, 0.87)"
                />
                <Text style={styles.regular25Out}>18</Text>
              </View>
              <Text style={styles.regular14}>check out</Text>
            </View>
          </View>
          <TouchableOpacity style={styles.addRecordStyle}>
            <Text style={styles.regular14Add}>Add record</Text>
            <Ionicons
              name="add-circle-outline"
              size={25}
              color="rgba(27, 97, 181, 0.89)"
            />
          </TouchableOpacity>
        </View>
        <View style={styles.detailContainer}>
          {/* <FlatList
            data={events}
            renderItem={renderItem}
            keyExtractor={(item) => item.id.toString()}
            contentContainerStyle={styles.eventList}
          /> */}
        </View>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#FFFFFF",
  },
  header_container: {},
  content_container: {
    flex: 1,
  },
  calendarContainer: {
    flex: 0.12,
    justifyContent: "center",
    backgroundColor: "blue",
  },
  calendarStripStyle: {
    height: 100,
  },
  statusContainer: {
    flex: 0.19,
    marginHorizontal: screenWidth * 0.05,
  },
  statusStyle: {
    flex: 0.65,
    flexDirection: "row",
    borderRadius: 5,
    justifyContent: "center",
    backgroundColor: "lightblue",
  },
  checkIn: {
    flex: 0.5,
    alignItems: "center",
    justifyContent: "center",
  },
  verticalLine: {
    height: "80%",
    width: 1,
    alignSelf: "center",
    backgroundColor: "rgba(108, 108, 108, 0.89)",
  },
  checkOut: {
    flex: 0.5,
    alignItems: "center",
    justifyContent: "center",
  },
  checkNum: {
    flex: 0.55,
    flexDirection: "row",
  },
  addRecordStyle: {
    flex: 0.35,
    alignItems: "center",
    justifyContent: "center",
    flexDirection: "row",
    marginTop: screenHeight * 0.005,
    borderBottomLeftRadius: 5,
    borderBottomRightRadius: 5,
    backgroundColor: "orange",
  },
  detailContainer: {
    flex: 0.69,
    backgroundColor: "yellow",
  },
  regular14: {
    fontSize: 14,
    color: "rgba(108, 108, 108, 0.89)",
    fontFamily: "Poppins_400Regular",
  },
  regular25In: {
    fontSize: 25,
    color: "rgba(52, 168, 83, 0.93)",
    fontFamily: "Poppins_400Regular",
    marginLeft: 10,
  },
  regular25Out: {
    fontSize: 25,
    color: "rgba(237, 74, 74, 0.87)",
    fontFamily: "Poppins_400Regular",
    marginLeft: 10,
  },
  regular14Add: {
    fontSize: 14,
    color: "rgba(27, 97, 181, 0.89)",
    fontFamily: "Poppins_400Regular",
    marginRight: 5,
  },
});
