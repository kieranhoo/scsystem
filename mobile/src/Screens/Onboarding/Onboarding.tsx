import { i18n, LocalizationKey } from "@/Localization";
import React from "react";
import {
  Dimensions,
  View,
  Text,
  StyleSheet,
  SafeAreaView,
  Image,
  TouchableOpacity,
  FlatList,
} from "react-native";
import { StatusBar } from "expo-status-bar";
import { HStack, Spinner, Heading } from "native-base";
import { User } from "@/Services";
import { Colors } from "../../Theme/Variables";
import { RootScreens } from "..";

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

const slides = [
  {
    id: "1",
    image: require("../../../assets/onboarding/illustration1.png"),
    title: "Scan Fast",
    subtitle:
      "User can use camera to scan student card for checkin or checkout faster than usual",
  },
  {
    id: "2",
    image: require("../../../assets/onboarding/illustration2.png"),
    title: "Monitor Easily",
    subtitle:
      "With history and dashboard, user can monitor all entry or exit activities of student easier",
  },
  {
    id: "3",
    image: require("../../../assets/onboarding/illustration3.png"),
    title: "Analyse Clearly",
    subtitle:
      "User can view all chart to analyse all entry activities of students more detail",
  },
];

const Slide: React.FC<{
  item: { id: string; image: number; title: string; subtitle: string };
}> = ({ item }) => {
  return (
    <View style={{ alignItems: "center" }}>
      <Image source={item?.image} style={styles.image} />
      <View>
        <Text style={styles.title}>{item?.title}</Text>
        <Text style={styles.subtitle}>{item?.subtitle}</Text>
      </View>
    </View>
  );
};

export const Onboarding = (props: { onNavigate: (string: RootScreens) => void }) => {
  const [currentSlideIndex, setCurrentSlideIndex] = React.useState(0);
  const ref = React.useRef<FlatList>(null);

  const updateCurrentSlideIndex = (e: any) => {
    const contentOffsetX = e.nativeEvent.contentOffset.x;
    const currentIndex = Math.round(contentOffsetX / screenWidth);
    setCurrentSlideIndex(currentIndex);
  };

  const goToNextSlide = () => {
    const nextSlideIndex = currentSlideIndex + 1;
    if (nextSlideIndex != slides.length) {
      const offset = nextSlideIndex * screenWidth;
      ref?.current?.scrollToOffset({ offset });
      setCurrentSlideIndex(currentSlideIndex + 1);
    }
  };

  const skip = () => {
    const lastSlideIndex = slides.length - 1;
    const offset = lastSlideIndex * screenWidth;
    ref?.current?.scrollToOffset({ offset });
    setCurrentSlideIndex(lastSlideIndex);
  };

  const Footer = () => {
    return (
      <View
        style={{
          flex: 1,
          justifyContent: "space-between",
          paddingHorizontal: 20,
        }}
      >
        <View style={styles.indicatorContainer}>
          {slides.map((_, index) => (
            <View
              key={index}
              style={[
                styles.indicator,
                currentSlideIndex == index && {
                  backgroundColor: Colors.DISPLAY,
                },
              ]}
            ></View>
          ))}
        </View>
        <View style={styles.buttonContainer}>
          {currentSlideIndex == slides.length - 1 ? (
            <View style={styles.nextButtonContainer}>
              <TouchableOpacity
                style={styles.btn}
                onPress={() => props.onNavigate(RootScreens.MAIN)}
              >
                <Text
                  style={{
                    fontFamily: "Poppins_700Bold",
                    fontSize: 21,
                    color: Colors.WHITE,
                  }}
                >
                  Get Started
                </Text>
              </TouchableOpacity>
            </View>
          ) : (
            <View style={{ flex: 1, flexDirection: "column" }}>
              <View style={styles.nextButtonContainer}>
                <TouchableOpacity
                  activeOpacity={0.8}
                  onPress={goToNextSlide}
                  style={styles.btn}
                >
                  <Text
                    style={{
                      fontFamily: "Poppins_700Bold",
                      fontSize: 21,
                      color: Colors.WHITE,
                    }}
                  >
                    Next
                  </Text>
                </TouchableOpacity>
              </View>
              <View style={styles.skipButtonContainer}>
                <TouchableOpacity
                  style={{ alignItems: "center" }}
                  onPress={skip}
                >
                  <Text
                    style={{
                      fontFamily: "Poppins_400Regular",
                      fontSize: 15,
                      textDecorationLine: "underline",
                      color: Colors.BLACK,
                    }}
                  >
                    Skip
                  </Text>
                </TouchableOpacity>
              </View>
            </View>
          )}
        </View>
      </View>
    );
  };

  return (
    <SafeAreaView style={styles.container}>
      <StatusBar style="auto" />
      <View style={styles.logoContainer}>
        <Image
          style={styles.logo}
          source={{
            uri: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyuA0nvDsTQ1oEs0UynwRL78ZVByg81UShUJJY9gkjZKauGaUV7sdPUqBqwT9QVymQ8Q&usqp=CAU",
          }}
        />
      </View>
      <View style={styles.bodyContainer}>
        <View style={styles.flatlistContainer}>
          <FlatList
            ref={ref}
            onMomentumScrollEnd={updateCurrentSlideIndex}
            data={slides}
            horizontal
            pagingEnabled
            showsHorizontalScrollIndicator={false}
            renderItem={({ item }) => <Slide item={item} />}
          />
        </View>
        <View style={styles.footerContainer}>
          <Footer />
        </View>
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: Colors.WHITE,
  },
  logoContainer: {
    flex: 0.2,
    alignItems: "center",
    justifyContent: "flex-end",
  },
  bodyContainer: {
    flex: 0.8,
  },
  flatlistContainer: {
    flex: 0.6,
  },
  footerContainer: {
    flex: 0.4,
  },
  logo: {
    width: 0.7 * screenWidth,
    height: 0.1 * screenHeight,
    resizeMode: "contain",
  },
  image: {
    height: "50%",
    width: screenWidth,
    resizeMode: "contain",
    marginVertical: 0.03 * screenHeight,
  },
  title: {
    fontSize: 28,
    textAlign: "center",
    fontFamily: "Poppins_700Bold",
  },
  subtitle: {
    fontSize: 17,
    color: Colors.GREY,
    textAlign: "center",
    fontFamily: "Poppins_400Regular",
    maxWidth: 0.7 * screenWidth,
  },
  indicatorContainer: {
    flex: 0.2,
    flexDirection: "row",
    justifyContent: "center",
    marginTop: 0.01 * screenHeight,
  },
  indicator: {
    height: 0.0175 * screenHeight,
    width: 0.0175 * screenHeight,
    borderRadius: (0.0175 * screenHeight) / 2,
    marginHorizontal: 0.03 * screenWidth,
    backgroundColor: Colors.HIDE,
  },
  buttonContainer: {
    flex: 0.8,
    flexDirection: "column",
    paddingHorizontal: 0.03 * screenWidth,
  },
  nextButtonContainer: {
    flex: 0.6,
    justifyContent: "flex-end",
    marginBottom: 0.02 * screenHeight,
  },
  skipButtonContainer: {
    flex: 0.4,
  },
  btn: {
    height: 0.08 * screenHeight,
    borderRadius: 12,
    backgroundColor: Colors.DISPLAY,
    justifyContent: "center",
    alignItems: "center",
  },
});
